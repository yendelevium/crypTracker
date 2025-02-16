import { Link } from "react-router"
function Navbar(){
    return(
        <>
        <nav className="flex justify-between p-3 bg-violet-400 text-xl text-white">
            <div>
                <Link to="/">crypTracker</Link>
            </div>
            <div className="navLinks flex justify-evenly">
                <Link to="cryptocurrencies">Currencies</Link>
                {/* Watchlist must redirect to login if  you're not logged in*/}
                <Link to="watchlist/<userId>">Watchlist</Link>
                {/* Either Login OR Profile must be displayed based on whether user is logged in or not */}
                <Link to="login">Login</Link>
                <Link to="signup">Signup</Link>
                <Link to="profile">Profile</Link>
            </div>
        </nav>
        </>
    )
}

export default Navbar