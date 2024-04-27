import './App.css';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Home from './components/Home';
import Login from './components/login/Login';
import Header from './components/navbar/Header';
import Orders from './components/order/Orders';
import Checkout from './components/checkout/Checkout';
import Payment from './components/order/Payment';

function App() {
  return (
    <Router>
      <div className="app">
        <Header /> {/* If Header should appear on all pages, place it outside of Routes but inside the Router */}
        <Routes>
          <Route path='/login' element={<Login />} />
          <Route path='/orders' element={<Orders />} />
          <Route path='/checkout' element={<Checkout />} />
          <Route path='/payment' element={<Payment />} />
          <Route path='/' element={<Home />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
