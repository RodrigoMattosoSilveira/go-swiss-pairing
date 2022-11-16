import React from 'react';
import { Member } from './Member';
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
                        <tr>
                            <td data-label="ID">{member.id}</td>
                            <td data-label="First">{member.first}</td>
                            <td data-label="Email">{member.email}</td>
                            <td data-label="Active">{member.isActive ? "Yes" : "No"}</td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </>
    )

}

export default MemberList;