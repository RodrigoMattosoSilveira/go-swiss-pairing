import { Member } from './Member';
import React from 'react';

interface ProjectCardProps {
    member: Member;
}

function MemberCard(props: ProjectCardProps) {
    const { member: member } = props;
    return (
        <tr key={member.id}>
            <td data-label="ID">{member.id}</td>
            <td data-label="Name">{member.first + " " + member.last}</td>
            <td data-label="Email">{member.email}</td>
            <td data-label="Cell">{member.cell}</td>
            <td data-label="Rating">{member.rating}</td>
            <td data-label="Active">{member.isActive ? "Yes" : "No"}</td>
        </tr>
    );
}

export default MemberCard;