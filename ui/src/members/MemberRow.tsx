import React from "react";
import { Member } from "./Member";

function formatDescription(description: string): string {
    return description.substring(0, 60) + '...';
}

interface MemberRowProps {
    member: Member;
}

function MemberRow (props: MemberRowProps) {
    const { member } = props;

    return (
        <tr>
            <td data-label="ID">{member.id}</td>
            <td data-label="First">{member.first}</td>
            <td data-label="Email">{member.email}</td>
            <td data-label="Active">{member.isActive ? "Yes" : "No"}</td>
        </tr>
    )
}

export default MemberRow