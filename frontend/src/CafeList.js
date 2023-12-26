
import React, { useState, useEffect } from 'react';
import Table from 'react-bootstrap/Table'
import axios from 'axios';



const CafeList = () => {
    const [products, setProducts] = useState([]);

    const readData = async () => {
        try {
            const response = await axios.get(`${window.api_location}/products`);
            console.log(response.data);
            setProducts(response.data);
        } catch (error) {
            console.log(error);
        }
    };

    useEffect(() => {
        readData();
    }, []);

    const getProducts = () => products.map((product, index) => (
        <tr key={index}>
            <td>{product.name}</td>
            <td>{product.price}</td>
            <td>{product.sku}</td>
        </tr>
    ));

    return (
        <div>
            <h1 style={{ marginBottom: "40px" }}>Menu</h1>
            <Table>
                <thead>
                    <tr>
                        <th>Name</th>
                        <th>Price</th>
                        <th>SKU</th>
                    </tr>
                </thead>
                <tbody>
                    {getProducts()}
                </tbody>
            </Table>
        </div>
    );
};

export default CafeList;