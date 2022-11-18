import React from 'react';
import {MOCK_MEMBERS} from "./MOCK_MEMBERS";
import MemberList from "./MemberList";
import MemberListTable from "./MemberListTable";

function MembersPage() {
    return (
        <>
            <h1>Members</h1>
            {/*<pre>{JSON.stringify(MOCK_MEMBERS, null, ' ')}</pre>*/}
            <MemberList members={MOCK_MEMBERS} />
            {/*<MemberListTable members={MOCK_MEMBERS} />*/}
        </>
        );
}

export default MembersPage;