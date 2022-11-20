import React from 'react';
import { Member } from './Member';

interface MemberListRowProps {
    member: Member;
    onEdit: (member: Member) => void;
}

function MemberListRow(props: MemberListRowProps) {
    // return <pre>{JSON.stringify(members, null, ' ')}</pre>;
    const { member, onEdit } = props;
    const handleEditClick = (memberBeingEdited: Member) => {
        // console.log(memberBeingEdited);
        onEdit(memberBeingEdited);
    };
    return (
        <>
            <div className="col-sm">
                <button
                    className="small"
                    onClick={() => {
                        handleEditClick(member);
                    }}
                >
                    <span className="icon-edit"></span>
                </button>
            </div>
            <div className="col-sm">{member.id}</div>
            <div className="col-sm">{member.first}</div>
            <div className="col-sm">{member.last}</div>
            <div className="col-sm">{member.email}</div>
            <div className="col-sm">{member.cell}</div>
            <div className="col-sm">{member.password}</div>
            <div className="col-sm">{member.rating}</div>
            <div className="col-sm">{member.isActive ? "Yes" : "No"}</div>
        </>
    )
}

export default MemberListRow;