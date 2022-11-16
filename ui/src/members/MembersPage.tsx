import React from 'react';
import {MOCK_MEMBERS} from "./MOCK_MEMBERS";

function MembersPage() {
    return (
        <>
            <h1>Members</h1>;
            <pre>{JSON.stringify(MOCK_MEMBERS, null, ' ')}</pre>
        </>
    )
}

export default MembersPage;