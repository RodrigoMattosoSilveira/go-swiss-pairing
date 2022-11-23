import React from 'react';

function MemberListRowHeader() {
    return (
        <>
            <div className="col-sm-1">Edit</div>
            <div className="col-sm-1">ID</div>
            <div className="col-sm-1">First</div>
            <div className="col-sm-1">Last</div>
            <div className="col-sm-2">Email</div>
            <div className="col-sm-2">Cell</div>
            <div className="col-sm-2">Password</div>
            <div className="col-sm-1">Rating</div>
            <div className="col-sm-1">Active</div>
        </>
    )
}

export default MemberListRowHeader;
