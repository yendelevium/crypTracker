import { TCoin } from "../utils/types"
import userStore from "../store/userStore"
import { useNavigate } from "react-router"
import toastStore from "../store/toastStore"

type CoinProps = {
    coinData : TCoin
}

function Coin(props:CoinProps){
    const {currentUser, watchlist, setWatchlist} = userStore()
    const {setToastMessage, setToastType} = toastStore();
    const navigator = useNavigate();
    // Checking to see if the coin is in the watchlist or not for conditional rendering
    function inWatchlist(targetCoin: TCoin) {
        // Not using for-each as ot doesn't support returns? TF
        for (let index = 0; index < watchlist.length; index++) {
            const coin = watchlist[index];
            if(coin.id === targetCoin.id){
                return true
            }
        }
        return false
    }

    // Adding coin to watchlist
    async function handleAdd(){
        const coinID = props.coinData.id
        const userID = currentUser?.user_id
        if (userID===undefined){
            setToastMessage(`Login to add ${coinID} to your watchlist!`)
            setToastType("error")
            navigator("/login")
            return
        }
        fetch(`/users/${userID}/watchlist`,{
            method: "POST",
            body: JSON.stringify({
                "user_id":userID,
                "coin_id":coinID
            })
        })
        .then(async (coinData)=>{
            const responseBody = await coinData.json()
            // If duplicate key, toast it lol
            if (coinData.status===500){
                setToastMessage(responseBody.message)
                setToastType("error")
                return
            }
            setToastMessage(`Added ${props.coinData.name} to watchlist!`)
            setToastType("success")
            setWatchlist([...watchlist, props.coinData])
        })
        .catch((error:unknown)=>{
            if (typeof error === "string") {
                setToastMessage(error)
                setToastType("error")
            } else if (error instanceof Error) {
                setToastMessage(error.message)
                setToastType("error")
            }
        })
    }

    async function handleRemove(){
        const coinID = props.coinData.id
        const userID = currentUser?.user_id

        // Making sure the user is loggedin before adding/removing the coin from the watchlist
        if (userID===undefined){
            setToastMessage(`Login to remove ${coinID} to your watchlist!`)
            setToastType("error")
            navigator("/login")
            return
        }
        fetch(`/users/${userID}/watchlist`,{
            method: "DELETE",
            body: JSON.stringify({
                "user_id":userID,
                "coin_id":coinID
            })
        })
        .then(async (coinData)=>{
            const responseBody = await coinData.json()
            if (coinData.status===500){
                setToastMessage(responseBody.message)
                setToastType("error")
                return
            }
            setToastMessage(`Removed ${props.coinData.name} to watchlist!`)
            setToastType("success")
            setWatchlist(watchlist.filter(coin =>{
                if (coin.id!==props.coinData.id){
                    return coin
                }
            }))
        })
        .catch((error:unknown)=>{
            if (typeof error === "string") {
                setToastMessage(error)
                setToastType("error")
            } else if (error instanceof Error) {
                setToastMessage(error.message)
                setToastType("error")
            }
        })
    }
    const currTime = new Date()
    
    return(
        <tr className="bg-white border-b border-gray-200 hover:bg-gray-50">
            <th scope="row" className="flex items-center px-6 py-4 text-gray-900 whitespace-nowrap dark:text-white">
                <img className="w-10 h-10 rounded-full" src={props.coinData.image} alt={`${props.coinData.id}'s Image`}/>
                <div className="ps-3">
                    <div className="text-base font-semibold text-gray-500">{props.coinData.id}</div>
                    <div className="font-normal text-gray-500">{props.coinData.name}</div>
                </div>  
            </th>
            <td className="px-6 py-4">
                {props.coinData.current_price} $
            </td>
            <td className="px-6 py-4">
                {currTime.toUTCString()}
            </td>
            <td className="px-6 py-4">
                {inWatchlist(props.coinData)===true?
                    <button className="cursor-pointer" onClick={handleRemove}>Remove from Watchlist</button>
                    :<button className="cursor-pointer" onClick={handleAdd}>Add to Watchlist</button>
                }
            </td>
        </tr>
    )
}

export default Coin