import React from 'react';
import './Home.css'; // Ensure the path is correct based on your folder structure
import { products } from '../mock/products';

function Home() {
    return (
        <div>
            <h1>Products</h1>
            <div className="products-container">
                {products.map((product) => (
                    <div key={product.productId} className="product-card">
                        <h2 className="product-title">{product.title}</h2>
                        <img src={product.url} alt={product.name} className="product-image" />
                        <p className="product-description">{product.description}</p>
                        <p className="product-price">Price: ${product.price}</p>
                    </div>
                ))}
            </div>
        </div>
    );
}

export default Home;