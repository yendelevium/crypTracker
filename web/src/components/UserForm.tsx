import React from "react"
import { useNavigate } from 'react-router';
import Toast from "./Toast";
import userStore from "../store/userStore";

export default function UserForm(props?:any){
    const navigate = useNavigate();
    const {setToastMessage, toastMessage} = userStore()
    // Making the form a controlled component by controlling the state of the inputs
    const [credentials,setCredentials] = React.useState({
        username:"",
        password:""
    })

    function handleChange(event: React.ChangeEvent<HTMLInputElement>){
        // Getting the name of the input field and updating it's value in state
        // This is necessary because name is a variable, not a string literal. If you wrote:
        // It would set the key "name" instead of using the actual value stored in name
        const {value, name}:{value: string, name: string} = event.currentTarget
        setCredentials(prevState=>{
            // Idk why we need [name], just 'name' wasn't working
            return {...prevState, [name]:value}
        })
    }

    function handleSubmit(event: React.FormEvent<HTMLFormElement>){
        // Preventing refresh and resetting all fields to empty on submit
        event.preventDefault()
        // If the login is successful, go to the prev page u were on, otherwise create a toast and STAY on the
        // login/signup page
        props.handleAuth(credentials.username,credentials.password)
            .then((data:any)=>{
                // We can't pass state to the previous using navigate(-1)
                // So, we use the global store and set the toastMessage there instead
                setToastMessage(data.message)
                navigate(-1)
            })
            .catch((error:unknown)=>{
                console.log(error)
                if (typeof error === "string") {
                    setToastMessage(error)
                } else if (error instanceof Error) {
                    setToastMessage(error.message)
                }
            })
        // Toast to show whether login was successfull or not
        setCredentials({username:"",password:""})
    }

    return(
        <div>
            {toastMessage ?<Toast message={toastMessage} type={"error"}/> : null}
            <form className="max-w-sm mx-auto" onSubmit={handleSubmit}>
                <div className="mb-5">
                    <label htmlFor="username" className="block mb-2 text-sm font-medium text-gray-900">Username</label>
                    <input type="text" id="username" name="username" onChange={handleChange} value={credentials.username} className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5" placeholder="John Doe" required />
                </div>
                <div className="mb-5">
                    <label htmlFor="password" className="block mb-2 text-sm font-medium text-gray-900">Password</label>
                    <input type="password" id="password" name="password" onChange={handleChange} value={credentials.password} className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5" required />
                </div>
                <button type="submit" className="w-full text-white bg-purple-700 hover:bg-purple-800 focus:outline-none focus:ring-4 focus:ring-purple-300 font-medium rounded-full text-sm px-5 py-2.5 text-center mb-2 dark:bg-purple-600 dark:hover:bg-purple-700 dark:focus:ring-purple-900 ">Submit</button>
            </form>
        </div>
    )
} 