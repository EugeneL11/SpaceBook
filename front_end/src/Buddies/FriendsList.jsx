import { React, useState,useEffect } from "react";

function Friend(props) {
    const removeFriendEvent = () => {
        props.removeFriend(props.username)
    }
    const othersProfileEvent = props.toggleOtherProfile
    console.log(othersProfileEvent)
    return (
        <div className="flex flex-row bg-blue-500 h-20 w-1/2 justify-between rounded-md">
            {/* <div className="inline-block m-1">
                <img
                    className="rounded-lg w-20 h-20"
                    src={props.user_pic_url}
                    alt={props.username} 
                ></img>
                <p>{props.username}</p>
            </div> */}

            <div onClick = {othersProfileEvent} className="flex items-center">
                <img 
                    src={props.user_pic_url}
                    alt={props.username} 
                    className="w-12 h-12 rounded-full aspect-square p-2"
                ></img>
                <p className="text-lg">{props.username}</p>
            </div>

            <div 
                onClick={removeFriendEvent}
                className="w-24 flex flex-col hover:-translate-y-1 transition-all ease-in-out items-center justify-center cursor-pointer 
                text-black hover:after:text-xs hover:after:text-center hover:after:content-['Unorbit']"
            >
                <img src="./remove-friend.png" className="w-8 h-8"/>
            </div>

            {/* <div className="inline-block m-1" onClick={removeFriendEvent}>
                <img 
                    className="sr-only"
                    src={UNORBIT_IMG_URL}
                    alt="unorbit"
                ></img>
                <p>unorbit</p>
            </div> */}

            {/* <div className="inline-block m-1">
                <img
                    className="sr-only"
                    src={WORMHOLE_IMG_URL}
                    alt="launch wormhole"
                ></img>
                <p>launch wormhole</p>
            </div> */}
        </div>
    );
}

function FriendsList(props) {
    const [friends, setFriends] = useState(null)

    const toggleMyProfile = props.toggleMyProfile
    const toggleFriendsList = props.toggleFriendsList
    const toggleOtherProfile = props.toggleOtherProfile
    useEffect(() =>{
        const friendstest = [{
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

        // ask back-end for friends list
        setFriends(friendstest)
    },[])
    const removeFriend = (friendtoRemove) => {
        const newFriendlist = friends.filter(
            (friend) => friend.username !== friendtoRemove
        );
        setFriends(newFriendlist);
        // do back-end stuff
    }
    
    return (
        <div>
            <button onClick={toggleMyProfile}> Back</button>
            <div className="flex-col">
                {friends ? friends.map(
                    (friend, index) => (
                        <div key={index} className="flex flex-col items-center mb-4">
                            <Friend 
                                username={friend.username} 
                                user_pic_url={friend.user_pic_url}
                                removeFriend = {removeFriend}
                                toggleOtherProfile = {() => {toggleOtherProfile(friend.username,toggleFriendsList)}}
                            ></Friend>
                        </div>
                    )
                ) : null}
            </div>
        </div>
    );
}

export default FriendsList;