import React from 'react';
import { Member } from './Member';
import MemberRow from "./MemberRow";
import {MOCK_MEMBERS} from "./MOCK_MEMBERS";

interface MemberListProps {
    members: Member[];
}

function MemberList({ members }: MemberListProps) {
    // return <pre>{JSON.stringify(members, null, ' ')}</pre>;
    return (
        <>
            <table className="hoverable">
                <thead>
                    <tr>
                        <th>ID</th>
                        <th>First</th>
                        <th>Email</th>
                        <th>Active</th>
                    </tr>
                </thead>
                <tbody>
                    {members.map((member) => (
                        <MemberRow member={member}></MemberRow>
                    ))}
                </tbody>
            </table>
        </>
    )

}

export default MemberList;