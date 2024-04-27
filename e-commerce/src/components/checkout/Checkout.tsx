import React from 'react';
import './Checkout.css';
import { useSelector } from 'react-redux';
import CheckoutProduct from './CheckoutProduct';
import { selectBasket } from '../../store/shoppingcart/shoppingSlice';
import { selectUser } from '../../store/users/userSlice';
import Subtotal from '../order/Subtotal';

const Checkout = () => {
    const itemsBasket = useSelector(selectBasket);
    const user = useSelector(selectUser);

    return (
        <div className="checkout">
            <div className="checkout__left">
                <div>
                    {user ? <h3>Greetings, {user.displayName} </h3> : null}
                    <h2 className="checkout__title" >Your Shopping Basket</h2>
                    {/* @ts-ignore */}
                    {itemsBasket.map(item => (
                        // @ts-ignore
                        <CheckoutProduct 
                          id={item.id}
                          title={item.title}
                          image={item.image}
                          price={item.price}
                         />
                    ))}
                </div>
            </div>
            <div className="checkout__right">
                <Subtotal />
            </div>
        </div>
    )
}

export default Checkout;
