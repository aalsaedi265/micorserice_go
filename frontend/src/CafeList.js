
import React from 'react';
import Table from 'react-bootstrap/Table'
import axios from 'axios';


//typically would use componemnts and its the right thing, but since it only two files
// using the old way won't make difference


class CafeList extends React.Component {

    getProduct() {
        let table = []
        
        return table
    }
    constructor(props) {
        super(props);
        this.readData();
        this.state = { products: [] };
    
        this.readData = this.readData.bind(this);
    }
    render() {
        return (
            <div>
                <h1 style={{ marginBottom: "40px" }}>Menu</h1>
                <Table>
                    <thead>
                        <tr>
                            <th>
                                Name
                            </th>
                            <th>
                                Price
                            </th>
                            <th>
                                SKU
                            </th>
                        </tr>
                    </thead>
                    <tbody>
                        {this.getProducts()}
                    </tbody>
                </Table>
            </div>
        )
    }
}

export default CafeList