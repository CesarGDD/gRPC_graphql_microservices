import React, { useEffect, useState } from 'react';
import { useSelector } from 'react-redux';
import { useGetOrdersDetailsByUserIdQuery, GetOrdersDetailsByUserIdQuery } from '../../graphql/generated/graphql'; // Import the types
import Order from './Order';
import './Orders.css';
import { selectUser } from '../../store/users/userSlice';

const Orders = () => {
    const user = useSelector(selectUser);

    // Explicitly type the orders state as an array of order details
    const [orders, setOrders] = useState<GetOrdersDetailsByUserIdQuery['getOrdersDetailsByUserId']>([]);

    // Use the query to fetch orders by user ID
    const [{ data, fetching, error }] = useGetOrdersDetailsByUserIdQuery({
        variables: { input: { userId: user?.userId } }, // Pass the user ID as the input
        pause: !user, // Pause the query if there's no user
    });

    // Update the orders state when data is fetched
    useEffect(() => {
        if (data && data.getOrdersDetailsByUserId) {
            setOrders(data.getOrdersDetailsByUserId);
        }
    }, [data]);

    if (fetching) return <p>Loading your orders...</p>;
    if (error) return <p>Failed to load orders. Please try again later.</p>;

    return (
        <div className="orders">
            <h1>Your orders</h1>
            <div className="orders__order">
                {orders.length === 0 ? (
                    <h2>YOU DON'T HAVE ANY ORDERS YET</h2>
                ) : (
                    orders.map(order => <Order key={order.orderId} order={order} />)
                )}
            </div>
        </div>
    );
};

export default Orders;
