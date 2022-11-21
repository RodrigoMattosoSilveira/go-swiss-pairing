import React from 'react';
import { Member } from './Member';

interface MemberDetailProps {
    member: Member;
}
export default function MemberDetail({ member }: MemberDetailProps) {
    return (
        <div className="row">
            <div className="col-sm-6">
                <div className="card large">
                    <img width="49" height="99"
                        className="rounded"
                        src={member.imageUrl}
                        alt={member.first + " " + member.last}
                    />
                    <section className="section dark">
                        <h3 className="strong">
                            <strong>{member.first + " " + member.last}</strong>
                        </h3>
                        <p>Email:  {member.email}</p>
                        <p>cell:   {member.cell}</p>
                        <p>rating: {member.cell}</p>
                        <p>
                            <mark className="active">
                                {' '}
                                {member.isActive ? 'active' : 'inactive'}
                            </mark>
                        </p>
                    </section>
                </div>
            </div>
        </div>
    );
}