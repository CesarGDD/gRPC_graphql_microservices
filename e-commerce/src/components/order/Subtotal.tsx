import React from 'react';
import './Subtotal.css';
import CurrencyFormat from 'react-currency-format';
import { useSelector } from 'react-redux';
import { getBasketTotal, selectBasket } from '../../store/shoppingcart/shoppingSlice';
import { selectUser } from '../../store/users/userSlice';


const Subtotal = () => {
    const itemsBasket = useSelector(selectBasket);
    const user = useSelector(selectUser);
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
            {/* <button onClick={!user ? e => history.push('/login') : e => history.push('/payment')} >Proceed to Checkout</button> */}
        </div>
    )
}

export default Subtotal;
