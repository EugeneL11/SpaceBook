import { React, useState, useEffect } from "react";
import {userID} from "../Static.js"


function Request(props) {
    // const removeFriendEvent = () => {
    //     props.removeFriend(props.username)
    // }
    const othersProfileEvent = props.toggleOtherProfile
    const denyRequestEvent = () => {
        props.denyRequest(props.username)
    }
    const acceptRequestEvent = () => {
        props.acceptRequest(props.username)
    }

    console.log(othersProfileEvent)
    return (
        <div className="flex flex-row bg-blue-500 hover:bg-blue-400 h-20 w-1/2 justify-between rounded-md px-5">
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

        // ask back-end for orbit requests
        setRequests(exampleRequests)
    },[])

    const denyRequest = (requestToDeny) => {
        const newRequestList = requests.filter(
            (request) => request.username !== requestToDeny
        )
        setRequests(newRequestList);
        // do back-end stuff
    }

    const acceptRequest = (requestToAccept) => {
        const newRequestList = requests.filter(
            (request) => request.username !== requestToAccept
        )
        setRequests(newRequestList);
        // do back-end stuff
    }

    return (<div className="flex flex-col">
        <div className="flex flex-start w-full">
            <button className="mb-5 w-fit ml-6 text-3xl hover:text-purple-300" onClick={toggleHomepage}> {'←'} </button>
        </div>
        {/* {exampleRequests.map((friend) =>(
            <button onClick={() => {toggleOtherProfile(friend, toggleNotifications)}}> See Other Profile: {friend}</button>
        ))} */}

        {requests ? requests.map(
            (request, index) => (
                <div key={index} className="flex flex-col items-center mb-8">
                    <Request 
                        username={request.username} 
                        user_pic_url={request.user_pic_url}
                        denyRequest = {denyRequest}
                        acceptRequest = {acceptRequest}
                        toggleOtherProfile = {() => {toggleOtherProfile(request.username, toggleNotifications)}}
                    ></Request>
                </div>
            )
        ) : null}
    </div>);

}

export default Notifications;