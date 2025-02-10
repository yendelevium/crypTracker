import { Link } from "react-router"
function Login(){
    return(
        <>
            <h1>Login</h1>
            <div>New User? <Link to="/signup">Signup</Link></div>
        </>
    )
}

export default Login