import { Link } from "react-router"
import UserForm from "../components/UserForm"
function Login(){
    return(
        <main className="p-3">
            <h1 className="text-4xl pb-2 max-w-sm mx-auto text-center">Login</h1>
            <UserForm />
            <div className="max-w-sm mx-auto text-center">New User? <Link to="/signup">Signup</Link></div>
        </main>            
    )
}

export default Login