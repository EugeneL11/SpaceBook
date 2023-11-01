import { React, useState,useEffect } from "react";

function Friend(props) {
    const UNORBIT_IMG_URL = ""
    const WORMHOLE_IMG_URL = ""
    const removeFriendEvent = () => {props.removeFriend(props.username)}
    return (
        <div class="flex-row">
            <div class="inline-block m-1">
                <img
                    class="rounded-lg w-20 h-20 sr-only"
                    src={props.user_pic_url}
                    alt={props.username} 
                ></img>
                <p>{props.username}</p>
            </div>
            <div class="unorbit-btn inline-block m-1" onClick={removeFriendEvent}>
                <img 
                    class="sr-only"
                    src={UNORBIT_IMG_URL}
                    alt="unorbit"
                ></img>
                <p>unorbit</p>
            </div>
            <div class="launch-wormhole-btn inline-block m-1">
                <img
                    class="sr-only"
                    src={WORMHOLE_IMG_URL}
                    alt="launch wormhole"
                ></img>
                <p>launch wormhole</p>
            </div>
        </div>
    );
}

function FriendsList(props) {
    const [friends, setFriends] = useState(null)

    useEffect(() =>{
        const friendstest = [{
        username: "Gene",
        user_pic_url: "/assets/user0-pfp.jpg"
    }] // placeholder for back-end data

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
        <div class="flex-col">
            {friends ? friends.map(
                (friend, index) => (
                    <div>
                        <Friend 
                            key={index}
                            username={friend.username} 
                            user_pic_url={friend.user_pic_url}
                            removeFriend = {removeFriend}
                        ></Friend>
                    </div>
                )
            ) : null}
        </div>
    );
}

export default FriendsList;