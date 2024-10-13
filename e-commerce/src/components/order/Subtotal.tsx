import React from 'react';
import './Subtotal.css';
import CurrencyFormat from 'react-currency-format';
import { useDispatch, useSelector } from 'react-redux';
import { cleanBasket, getBasketTotal, selectBasket } from '../../store/shoppingcart/shoppingSlice';
import { selectUser } from '../../store/users/userSlice';
import { NewOrderItem, useCreateOrderMutation } from '../../graphql/generated/graphql'; // Import the mutation hook
import { useNavigate } from 'react-router-dom';

const Subtotal = () => {
    const itemsBasket = useSelector(selectBasket);
    const user = useSelector(selectUser);
    const navigate = useNavigate();
    const dispatch = useDispatch(); 
    const [{ fetching, error }, createOrder] = useCreateOrderMutation();

    const handleCheckout = async () => {
        if (!user) {
            navigate('/login');
            return;
        }
        
        // Explicitly define the type for the productMap and orderItems
        const productMap: { [key: number]: NewOrderItem } = itemsBasket.reduce((acc: any, item: any) => {
            if (acc[item.id]) {
                acc[item.id].quantity += 1;
            } else {
                acc[item.id] = { productId: item.id, quantity: 1 };
            }
            return acc;
        }, {} as { [key: number]: NewOrderItem });
    
        // Convert the product map to an array of orderItems (NewOrderItem[])
        const orderItems: NewOrderItem[] = Object.values(productMap);
    
        const totalPrice = getBasketTotal(itemsBasket);
    
        const input = {
            orderItems,
            totalPrice,
            userId: user.userId, // Assuming user object has an 'id' field
        };
    
        try {
            const response = await createOrder({ input });
    
            if (response.data?.createOrder.success) {
                alert('Order placed successfully!');
                dispatch(cleanBasket([])); 
                navigate('/orders');
            } else {
                alert('Failed to create the order.');
            }
        } catch (err) {
            console.error('Order creation failed:', err);
            alert('An error occurred. Please try again.');
        }
    };

    return (
        <div className="subtotal">
            <CurrencyFormat 
                renderText={(value) => (
                    <>
                        <p>
                          Subtotal ({itemsBasket.length} items):
                          <strong> {` ${value}`} </strong>
                        </p>
                        <small className="subtotal__gift" >
                            <input type="checkbox" /> 
                            This order contains a gift
                        </small>
                    </>
                )}
                decimalScale={2}
                value={getBasketTotal(itemsBasket)}
                displayType={"text"}
                thousandSeparator={true}
                prefix={"$"}
            />

            {user ? (
                <button onClick={handleCheckout} disabled={fetching}>
                    {fetching ? 'Processing...' : 'Proceed to Checkout'}
                </button>
            ) : (
                <p className="login__message">You must log in to proceed with the checkout.</p>
            )}

            {error && <p style={{ color: 'red' }}>Error creating order. Please try again.</p>}
        </div>
    );
}

export default Subtotal;
