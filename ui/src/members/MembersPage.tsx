import React from 'react';
import {MOCK_MEMBERS} from "./MOCK_MEMBERS";
import MemberList from "./MemberList";

function MembersPage() {
    return (
        <>
            <h1>Members</h1>
            <MemberList members={MOCK_MEMBERS} />
        </>
    )
}

export default MembersPage;