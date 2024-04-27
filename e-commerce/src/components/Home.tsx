import React, { useState } from "react";
import "./Home.css";
import Product from "../components/products/Product";
import Modal from "react-modal";
import { mockProducts } from "../mock/products";
import { Product as ProductType, useGetProductsQuery } from "../graphql/generated/graphql";


function chunkArray(array: any[], size: number): any[][] {
    const chunkedArr = [];
    for (let i = 0; i < array.length; i += size) {
      chunkedArr.push(array.slice(i, i + size));
    }
    return chunkedArr;
  }

const Home = () => {
  const [modalOpen, setModalOpen] = useState(true);

  const [{ data, fetching, error }] = useGetProductsQuery();
  let productsToDisplay: any = mockProducts; 

  if (data && data.getProducts) {
    productsToDisplay = data.getProducts;
  }

  const chunkedProducts = chunkArray(productsToDisplay, 4);

  const styleHandle = {
    style: {
      overlay: {
        backgroundColor: "rgba(17,17,17,0.9)",
      },
      content: {
        width: "60%",
        top: "50%",
        left: "50%",
        right: "auto",
        bottom: "auto",
        marginRight: "-50%",
        transform: "translate(-50%, -50%)",
        background: "rgba(52,58,64,0.9)",
      },
    },
  };

      if (fetching) return <p>Loading...</p>;

      if (error) return <p>Error: {error.message}</p>;

  return (
    <div className="home">
      <div className="home__container">
        <img
          className="home__image"
          src="https://images.unsplash.com/photo-1482775907821-a56ec43344fc?ixlib=rb-1.2.1&ixid=MXwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHw%3D&auto=format&fit=crop&w=1352&q=80"
          alt=""
        />
        <Modal isOpen={modalOpen} {...styleHandle}>
          <div className="modal">
            <h1>
              This is a demo web app so payments are disabled, you can check the
              code to see the payment method. Thanks for your understanding.
            </h1>

            <button onClick={() => setModalOpen(false)}>CLOSE</button>
          </div>
        </Modal>
        <div>
          {chunkedProducts.map((moreData, index) => (
            <div className="home__row" key={index}>
              {moreData.map((product: ProductType) => (
                <Product
                  key={product.productId}
                  title={product.title}
                  price={product.price}
                  id={product.productId}
                  image={product.url}
                />
              ))}
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};

export default Home;
