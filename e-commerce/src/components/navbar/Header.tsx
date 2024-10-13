import React from 'react';
import SearchIcon from '@material-ui/icons/Search';
import ShoppingBasketIcon from '@material-ui/icons/ShoppingBasket';
import './Header.css';
import { Link, useNavigate } from 'react-router-dom'; // Import useNavigate
import { useDispatch, useSelector } from 'react-redux';
import { selectUser, logout } from '../../store/users/userSlice';
import { selectBasket } from '../../store/shoppingcart/shoppingSlice';

const Header = () => {
    const user = useSelector(selectUser); // Get the user from Redux
    const dispatch = useDispatch();
    const navigate = useNavigate(); // Get the navigate function
    const itemsBasket = useSelector(selectBasket); // Get the basket items from Redux

    // Handle sign out
    const handleSignOut = () => {
        dispatch(logout());
        navigate('/'); // Redirect to home page after sign out
    };

    return (
        <div className="header">
            <Link className="header__link" to='/'>
                <h1 className="header__logo">E-commerce</h1>
            </Link>
            <div className="header__search">
                <input type="text" className="header__searchInput" />
                <SearchIcon className="header__searchIcon" />
            </div>
            <div className="header__nav">
                <div className="header__option">
                    <span className="header__optionLineOne">
                        {user ? `Hello, ${user.username}` : 'Greetings'}
                    </span>
                    {!user ? (
                        <Link to="/login" className="header__optionLineTwo">
                            <span>Sign In</span>
                        </Link>
                    ) : (
                        <span onClick={handleSignOut} className="header__optionLineTwo">
                            Sign Out
                        </span>
                    )}
                </div>
                <div className="header__option">
                    <Link className="header__link" to="/orders">
                        <span className="header__optionLineTwo">My Orders</span>
                    </Link>
                </div>
                <Link to="/checkout">
                    <div className="header__optionBasket">
                        <ShoppingBasketIcon />
                        <span className="header__optionLineTwo header__basketCount">
                            {itemsBasket?.length || 0}
                        </span>
                    </div>
                </Link>
            </div>
        </div>
    );
};

export default Header;
