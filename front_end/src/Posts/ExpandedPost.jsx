import { React, useState } from "react";

function ExpandedPost(props) {
    return (
        <div class="lg flex justify-center items-center relative border border-white border-solid bg-red-500">
            <button class="absolute top-0 left-0">&larr; Prev Page</button>
            <div class="md container">Hi</div>
        </div>
    );
}

export default ExpandedPost;
