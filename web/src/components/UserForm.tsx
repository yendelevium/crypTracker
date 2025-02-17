import React from "react"

export default function UserForm(props?:any){
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
        props.handleAuth(credentials.username,credentials.password)
        // Login/Signup logic ot backend
        // Toast to show whether login was successfull or not
        setCredentials({username:"",password:""})
    }

    return(
        <div>
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