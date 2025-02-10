import { Link } from "react-router"
function Signup(){
    return(
        <>
            <h1>Signup</h1>
            <div>Already registered? <Link to="/login">Login</Link></div>
        </>
    )
}

export default Signup