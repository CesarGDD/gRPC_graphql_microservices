import React from 'react';
import { useDispatch } from 'react-redux';
import './Product.css';
import { setBasket } from '../../store/shoppingcart/shoppingSlice';
import { ProductType } from '../../models/products';


const Product = ({title, image, price, id}: ProductType) => {
    const dispach = useDispatch()
    
    const addToBasket = () => {
        console.log("addToBasket Clicked")
        dispach(
            setBasket({
                        title: title,
                        image: image,
                        price: price,
                        id: id
            })
        )
    }

    return (
        <div className="product">
            <div className="product__info">
                <p> {title} </p>
                <p className="product__price">
                    $
                    <strong> {price} </strong> AUD
                </p>
                
            </div>
            <img src={image} alt=""/>
            <button onClick={addToBasket} >Add to Basket</button>
        </div>
    )
} 

export default Product;
