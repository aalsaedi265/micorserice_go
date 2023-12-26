
import React from 'react'
import Navbar from 'react-bootstrap/Navbar';
import './App.css';

import 'bootstrap/dist/css/bootstrap.min.css';


function App() {
  return (
    <div className="App">
      <Navbar bg='light' expand='lg'>
        <Navbar.Brand href='#home'>Cafe Dearly</Navbar.Brand>
        <Navbar.Toggle aria-controls="basic-navbar-nav" />

      </Navbar>
      
    </div>
  );
}

export default App;
