import { Link } from "react-router"
function Navbar(){
    return(
        <>
        <Link to="/">crypTracker</Link>
        <Link to="cryptocurrencies">Currencies</Link>

        {/* Watchlist must redirect to login if  you're not logged in*/}
        <Link to="watchlist/<userId>">Watchlist</Link>
        {/* Either Login OR Profile must be displayed based on whether user is logged in or not */}
        <Link to="login">Login</Link>
        <Link to="profile">Profile</Link>
        </>
    )
}

export default Navbar