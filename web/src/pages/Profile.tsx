import userStore from "../store/userStore"
function Profile(){
    const {currentUser} = userStore();
    console.log("Profile?",currentUser)
    return(
        <main className="p-3">
            <h1 className="text-4xl pb-2 text-center">User Profile:{currentUser?.username}</h1>
            <div className="flex justify-center">
                <div className="max-w-sm bg-white border border-gray-200 rounded-lg shadow-sm">
                    <img className="rounded-t-lg h-48 w-96 object-cover" src={currentUser?.profile_image} alt="profile-pic" />
                    <div className="p-5">
                        <h5 className="mb-2 text-2xl font-bold tracking-tight text-gray-900">{currentUser?.username}</h5>
                        <p className="mb-3 font-normal text-gray-700 dark:text-gray-400">Slay the Way</p>
                    </div>
                </div>
            </div>
        </main>
    )
}

export default Profile