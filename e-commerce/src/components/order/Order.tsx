import React from 'react';
import moment from 'moment';
import './Order.css';
import CheckoutProduct from '../checkout/CheckoutProduct';
import CurrencyFormat from 'react-currency-format';
import { useGetProductQuery } from '../../graphql/generated/graphql'; // Import product query

// @ts-ignore
const Order = ({ order }) => {
    return (
        <div className="order">
            <h2>Order</h2>
            <p>Order Date: {moment(order.createdAt).format('MMMM Do YYYY, h:mma')}</p>
            <p className="order__id">
                <small>Order ID: {order.orderId}</small>
            </p>
            
            {/* Render each product in the order */}
            {order.orderItems.map((item: any )=> (
                <ProductItem key={item.productId} productId={item.productId} quantity={item.quantity} />
            ))}

            {/* Display total price */}
            <CurrencyFormat
                renderText={(value) => (
                    <h3 className="order__total">Order Total: {order.totalPrice}</h3>
                )}
                decimalScale={2}
                value={order.totalPrice / 100} // Assuming totalPrice is in cents
                displayType={"text"}
                thousandSeparator={true}
                prefix={"$"}
            />
        </div>
    );
};

// New component for fetching and displaying product info
const ProductItem = ({ productId, quantity }: any) => {
    const [{ data, fetching, error }] = useGetProductQuery({
        variables: { productId },
    });

    if (fetching) return <p>Loading product details...</p>;
    if (error) return <p>Failed to load product details</p>;

    const product = data?.getProduct;

    return (
        <CheckoutProduct
            id={product?.productId}
            title={product?.title}
            image={product?.url || '/default-product-image.jpg'}
            price={product?.price}
            
            hideButton
        />
    );
};

export default Order;
