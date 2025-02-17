import { Link } from "react-router"
import UserForm from "../components/UserForm"
import userStore from "../store/userStore"

function Signup(){
    const { handleUserSignup }= userStore()
    return(
        <main className="p-3">
            <h1 className="text-4xl pb-2 max-w-sm mx-auto text-center">Signup</h1>
            <UserForm handleAuth={handleUserSignup}/>
            <div className="max-w-sm mx-auto text-center">Already registered? <Link to="/login">Login</Link></div>
        </main>   
    )
}

export default Signup