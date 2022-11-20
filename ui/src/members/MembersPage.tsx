import React, { Fragment, useState } from 'react';
import { MOCK_MEMBERS } from "../../src/members/MOCK_MEMBERS";
import MemberList from "../../src/members/MemberList"
import { Member } from "./Member"

function MembersPage() {
    const [members, setMembers] = useState<Member[]>(MOCK_MEMBERS);
    const saveMember = (member: Member) => {
        // console.log('Saving project: ', member);
        let updatedMembers = members.map((m: Member) => {
            return m.id === member.id ? member : m;
        });
        setMembers(updatedMembers);
    };
    return (
        <>
            <h1>Members</h1>
            {/* <pre>{JSON.stringify(MOCK_MEMBERS, null, ' ')}</pre> */}
            <MemberList members={members} onSave={saveMember}/>
        </>
    );
}

export default MembersPage;