import React from 'react'
import coinStore from '../store/coinStore'
import userStore from '../store/userStore'
import { TCoin } from '../utils/types'

// Table component to display coins in Currencies and Watchlist pages (and others, if any)
function Table(props:{coinJSX: React.JSX.Element[]}) {
    const {currentUser, setWatchlist} = userStore()
    const {allCoins } = coinStore()
    // Getting the watchlist of the user, and if it's null, setting it to [] coz otherwise JS will raise an error
    // Reget watchlist everytime my allCoins updates to be uptodate with the newest crypto prices
    // I'm getting the watchlist here instead of in /watchlist as upon signin/signout, I need to get the watchlist
    // in /cryptocurrencies, so that I can correctly display "add to watchlist" or "remove from watchlist"
    React.useEffect(()=>{
        const getWatchlist = async()=>{
            if (currentUser===null) return
            const fetchResponse = await fetch(`/users/${currentUser?.user_id}/watchlist`)
            let watchlistData: TCoin[] = await fetchResponse.json()
            if (watchlistData === null){
                watchlistData = []
            }
            setWatchlist(watchlistData)
        }
        getWatchlist()
    },[allCoins])

  return (
    <div className="relative overflow-x-auto shadow-md sm:rounded-lg">
        <table className="w-full text-sm text-left rtl:text-right text-gray-500">
            <thead className="text-xs text-gray-700 uppercase bg-gray-50">
                <tr>
                    <th scope="col" className="px-6 py-3">
                        Coin
                    </th>
                    <th scope="col" className="px-6 py-3">
                        Market Price
                    </th>
                    <th scope="col" className="px-6 py-3">
                        Last Updated At
                    </th>
                    <th scope="col" className="px-6 py-3">
                        Action
                    </th>
                </tr>
            </thead>
            <tbody>
                {props.coinJSX}
            </tbody>
        </table>
    </div>
  )
}

export default Table