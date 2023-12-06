import { React, useState, useEffect } from "react";
import currentUser from "../Static.js";
import {serverpath} from "../Path.js";

import axios from 'axios'
function Request(props) {
    // const removeFriendEvent = () => {
    //     props.removeFriend(props.username)
    // }
    const othersProfileEvent = props.toggleOtherProfile
    const denyRequestEvent = props.denyRequest
    const acceptRequestEvent = props.acceptRequest

    console.log(othersProfileEvent)
    return (
        <div className="flex flex-row bg-blue-500 hover:bg-blue-400 h-20 w-11/12 sm:w-3/4 lg:w-1/2 min-w-fit justify-between rounded-md px-5">
            <div onClick = {othersProfileEvent} className="flex items-center hover:cursor-pointer hover:opacity-70 hover:text-gray-700">
                <img 
                    src={props.user_pic_url}
                    alt={props.username} 
                    className="w-12 h-12 rounded-full aspect-square p-2"
                ></img>
                <p className="text-lg">{props.username}</p>
            </div>

            <div className="flex">
                <div 
                    onClick={denyRequestEvent}
                    className="w-12 flex flex-col hover:-translate-y-1 transition-all ease-in-out items-center justify-center cursor-pointer 
                    text-black hover:after:text-xs hover:after:text-center hover:after:content-['Deny']"
                >
                    <img src="./remove-friend.png" className="w-8 h-8"/>
                </div>

                <div 
                    onClick={acceptRequestEvent}
                    className="w-12 flex flex-col hover:-translate-y-1 transition-all ease-in-out items-center justify-center cursor-pointer 
                    text-black hover:after:text-xs hover:after:text-center hover:after:content-['Accept']"
                >
                    <img src="./accept-request.png" className="w-8 h-8"/>
                </div>
            </div>
            
        </div>
    );
}

function Notifications(props) {
    const [requests, setRequests] = useState(null)
    const [noReqs, setNoReqs] = useState(null)
    
    const toggleHomepage = props.toggleHomepage
    const toggleOtherProfile = props.toggleOtherProfile
    const toggleNotifications = props.toggleNotifications
    

    useEffect(() =>{
        const exampleRequests = [{
            username: "Gene",
            user_pic_url: "./jupiter.jpg"
        },
        {
            username: "Raine",
            user_pic_url: "./jupiter.jpg"
        },
        {
            username: "Omar",
            user_pic_url: "./jupiter.jpg"
        },
        {
            username: "Kevin",
            user_pic_url: "./jupiter.jpg"
        },

    ] // placeholder for back-end data

    //WE ARE GETTING THE FRIEND REQUESTS HERE, TO USE FOR DENYING/ACCEPTING (MUST FIX)
    const path = `/friendrequests/${encodeURIComponent(currentUser.userID)}`
    axios.post(`${serverpath}${path}`).then((res) => {
        const data = res.data
        console.log(data)
        if (data.status === "no requests") {
            setNoReqs(true)
        } else if (data.status === "pending request") {
            setNoReqs(false)
            setRequests(data.requests)
        } else {
            console.log("database error")
        }
    })

        // ask back-end for orbit requests
    },[])

    const denyRequest = (userToDeny) => {
        const path = `/rejectfriendreq/${encodeURIComponent(currentUser.userID)}/${encodeURIComponent(userToDeny)}`
        axios.delete(`${serverpath}${path}`).then((res) => {
            const data = res.data
            if (data.status === "no error") {
                const newRequestList = requests.filter(
                    (request) => request.user_id !== userToDeny
                )
                setRequests(newRequestList);
            } else {
                console.log(data.status)
            }
        })
    }

    const acceptRequest = (userToAccept) => {
        const path = `/sendfriendreq/${encodeURIComponent(currentUser.userID)}/${encodeURIComponent(userToAccept)}`
        axios.post(`${serverpath}${path}`).then((res) => {
            const data = res.data
            if (data.status === "no error") {
                const newRequestList = requests.filter(
                    (request) => request.user_id !== userToAccept
                )
                setRequests(newRequestList);
                console.log("accepted")
            } else {
                console.log(data.status)
            }
        })
    }

    return (<div className="flex flex-col">
        <div className="flex flex-start w-full">
            <button className="mb-2 w-fit ml-6 text-4xl hover:text-purple-300" onClick={toggleHomepage}> {'←'} </button>
        </div>
        {/* {exampleRequests.map((friend) =>(
            <button onClick={() => {toggleOtherProfile(friend, toggleNotifications)}}> See Other Profile: {friend}</button>
        ))} */}
        <h3 className="mx-auto mb-4 text-3xl text-white">Orbit Requests</h3>
        {noReqs ? 
            <div className="w-fit bg-white rounded-lg text-black text-center text-xl mx-auto p-10">
                No Orbit Requests... LOSER (jk)
            </div> 
        : null}
        {requests ? requests.map(
            (request, index) => (
                <div key={index} className="flex flex-col items-center mb-8">
                    <Request 
                        username={request.user_name} 
                        user_pic_url={serverpath + request.profile_picture_path}
                        denyRequest = {() => denyRequest(request.user_id)}
                        acceptRequest = {() => acceptRequest(request.user_id)}
                        toggleOtherProfile = {() => {toggleOtherProfile(request.user_id, toggleNotifications)}}
                    ></Request>
                </div>
            )
        ) : null}
    </div>);

}

export default Notifications;