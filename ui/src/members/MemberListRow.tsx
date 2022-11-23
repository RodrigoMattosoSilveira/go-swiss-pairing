import React from 'react';
import { Link } from 'react-router-dom';
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
        <div key={member.id} className="row">
            <div className="col-sm-1">
                <button
                    className="small"
                    onClick={() => {
                        handleEditClick(member);
                    }}
                >
                    <span className="icon-edit"></span>
                </button>
            </div>
            <Link to={'/members/' + member.id} className="col-sm-1">{member.id}</Link>
            <div className="col-sm-1">{member.first}</div>
            <div className="col-sm-1">{member.last}</div>
            <div className="col-sm-2">{member.email}</div>
            <div className="col-sm-2">{member.cell}</div>
            <div className="col-sm-2">{member.password}</div>
            <div className="col-sm-1">{member.rating}</div>
            <div className="col-sm-1">{member.isActive ? "Yes" : "No"}</div>
        </div>
    )
}

export default MemberListRow;