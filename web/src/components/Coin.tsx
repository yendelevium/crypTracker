import { TCoin } from "../utils/types"

type CoinProps = {
    coinData : TCoin
}

function Coin(props:CoinProps){
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
                <a href="#" className="font-medium text-blue-600 dark:text-blue-500 hover:underline">Add to Watchlist</a>
            </td>
        </tr>
    )
}

export default Coin