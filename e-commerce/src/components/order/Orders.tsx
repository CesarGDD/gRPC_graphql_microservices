import React, { useEffect, useState } from 'react';
import { useSelector } from 'react-redux';
import Order from './Order';
import './Orders.css';
import { selectUser } from '../../store/users/userSlice';

const Orders = () => {
    const user = useSelector(selectUser);
    const [orders, setOrders] = useState([]);

    // useEffect(() => {
    //     if (user) {
    //         db.collection('users')
    //           .doc(user?.uid)
    //           .collection('orders')
    //           .orderBy('created', 'desc')
    //           .onSnapshot(snapshot => (
    //               setOrders(snapshot.docs.map(doc => ({
    //                   id: doc.id,
    //                   data: doc.data()
    //               })))
    //           ))
    //     } else {
    //         setOrders([])
    //     }
    // },[user])

    return (
        <div className="orders" >
            <h1>Your orders</h1>
            <div className="orders__order">
                {orders.length === 0 ? <h2>YOU DONT'T HAVE ORDERS YET</h2> : null }
                {orders?.map(order => (
                    <Order order={order} />
                ))}
            </div>
        </div>
    )
}

export default Orders;