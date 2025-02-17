import userStore from "../store/userStore"
import coinStore from "../store/coinStore"
import Coin from "../components/Coin"
import { type TCoin } from "../utils/types"
import React from "react"
import Table from "../components/Table"

function Watchlist(){
    const {currentUser, watchlist, setWatchlist} = userStore()
    const {allCoins } = coinStore()
    let watchlistJSX: React.JSX.Element[] = [];

    // Getting the watchlist of the user, and if it's null, setting it to [] coz otherwise JS will raise an error
    // Reget watchlist everytime my allCoins updates to be uptodate with the newest crypto prices
    React.useEffect(()=>{
        const getWatchlist = async()=>{
            const fetchResponse = await fetch(`/users/${currentUser?.user_id}/watchlist`)
            let watchlistData: TCoin[] = await fetchResponse.json()
            if (watchlistData === null){
                watchlistData = []
            }
            setWatchlist(watchlistData)
        }
        getWatchlist()
    },[allCoins])

    // Only on a non-empty list (otherwise we get an error), map over the watchlist to create the JSX
    if (watchlist.length > 0){
        watchlistJSX = watchlist.map(coin=>{
            return (<Coin coinData={coin} key={coin.id}/>)
        })
    }

    return(
        <main className="p-3">
            <h1 className="text-4xl pb-2">Watchlist</h1>
            <div className="mb-2">
                The cryptocurrencies you are interested in
            </div>

            {/* Remove the 'Add to Watchlist' Thingy */}
            <Table coinJSX={watchlistJSX}/>
        </main>
    )
}

export default Watchlist