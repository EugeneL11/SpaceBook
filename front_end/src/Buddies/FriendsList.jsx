import { React, useState,useEffect } from "react";
import axios from 'axios'
import currentUser from "../Static";
import { serverpath } from "../Path";
function Friend(props) {
    const removeFriendEvent = props.removeFriend
    const othersProfileEvent = props.toggleOtherProfile

    //html code for a single friend
    return (
        <div className="flex flex-row bg-blue-500 hover:bg-blue-400 h-20 w-11/12 sm:w-3/4 lg:w-1/2 min-w-fit px-5 justify-between rounded-md">
            <div onClick = {othersProfileEvent} className="flex items-center hover:cursor-pointer hover:opacity-70 hover:text-gray-700">
                <img 
                    src={props.user_pic_url}
                    alt={props.username} 
                    className="w-12 h-12 rounded-full aspect-square p-2"
                ></img>
                <p className="text-lg">{props.username}</p>
            </div>

            <div 
                onClick={removeFriendEvent}
                className="w-16 flex flex-col hover:-translate-y-1 transition-all ease-in-out items-center justify-center cursor-pointer 
                text-black hover:after:text-xs hover:after:text-center hover:after:content-['Unorbit']"
            >
                <img src="./remove-friend.png" className="w-8 h-8"/>
            </div>
        </div>
    );
}

function FriendsList(props) {
    const [friends, setFriends] = useState(null)
    const toggleMyProfile = props.toggleMyProfile
    const toggleFriendsList = props.toggleFriendsList
    const toggleOtherProfile = props.toggleOtherProfile

    useEffect(() =>{
        // ask back-end for friends list
        const path = `/friends/${encodeURIComponent(currentUser.userID)}`
        axios.get(`${serverpath}${path}`).then((res) => {
            const data = res.data
            console.log(data)
            if (data.error === "no error") {
                setFriends(data.users)
            } else {
                console.log("ERROR")
            }
        })
    },[])

    const removeFriend = (friendToRemove) => {
        // tell the back-end to remove a user as a friend
        const path = `/removefriend/${encodeURIComponent(currentUser.userID)}/${encodeURIComponent(friendToRemove)}`
        axios.delete(`${serverpath}${path}`).then((res) => {
            const data = res.data
            if (data.status === "no error") {
                const newFriendlist = friends.filter(
                    (friend) => friend.user_id !== friendToRemove
                );
                setFriends(newFriendlist);
            } else {
                console.log(data.status)
            }
        })
    }
    
    // html code for viewing all of your friends in a list
    return (
        <div>
            <div className="flex flex-start w-full">
                <button className="mb-2 w-fit ml-6 text-3xl hover:text-purple-300" onClick={toggleMyProfile}> {'←'} </button>
            </div>
            <div className="flex-col">
                {friends ? friends.map(
                    (friend, index) => (
                        <div key={index} className="flex flex-col items-center mb-8">
                            <Friend 
                                username={friend.user_name} 
                                user_pic_url={serverpath + friend.profile_picture_path}
                                removeFriend = {() => removeFriend(friend.user_id)}
                                toggleOtherProfile = {() => {toggleOtherProfile(friend.user_id, toggleFriendsList)}}
                            ></Friend>
                        </div>
                    )
                ) : null}
            </div>
        </div>
    );
}

export default FriendsList;