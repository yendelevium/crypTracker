import React from 'react'

// Table component to display coins in Currencies and Watchlist pages (and others, if any)
function Table(props:{coinJSX: React.JSX.Element[]}) {
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
                        Last Updated AT (debugging purposes)
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