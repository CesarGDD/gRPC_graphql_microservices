import React, { useEffect, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import './Payment.css';
import { selectUser } from '../../store/users/userSlice';
import { selectBasket } from '../../store/shoppingcart/shoppingSlice';

const Payment = () => {
    const user = useSelector(selectUser);
    const itemsBasket = useSelector(selectBasket);
    const dispatch = useDispatch();

    useEffect(()=> {

    },[itemsBasket]);

    return (
        <>
        </>
    )
}

export default Payment;
