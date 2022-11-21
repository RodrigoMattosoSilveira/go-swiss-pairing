import React, { Fragment, useState, useEffect } from 'react';
import MemberList from "../../src/members/MemberList"
import { Member } from "./Member"
import { memberAPI } from "./memberAPI";

function MembersPage() {
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState<string | undefined>(undefined);
    const [members, setMembers] = useState<Member[]>([]);
    const [currentPage, setCurrentPage] = useState(1);
    useEffect(() => {
        async function loadMembers() {
            setMembers([]);
            try {
                setLoading(true);
                const data = await memberAPI.get(currentPage);
                setError('');
                setMembers(data);
                if (currentPage === 1) {
                    setMembers(data);
                } else {
                    setMembers((members) => [...members, ...data]);
                }
            }
            catch (e) {
                if (e instanceof Error) {
                    setError(e.message);
                }
            } finally {
                setLoading(false);
            }
        }
        loadMembers();
    }, [currentPage]);
    const saveMember = (member: Member) => {
        // console.log('Saving member: ', member);
        // let updatedMembers = members.map((m: Member) => {
        //     return m.id === member.id ? member : m;
        // });
        // setMembers(updatedMembers);
        memberAPI
          .put(member)
          .then((updatedMember) => {
                let updatedMembers = members.map((p: Member) => {
                      return p.id === member.id ? new Member(updatedMember) : p;
                    });
                setMembers(updatedMembers);
              })
          .catch((e) => {
                 if (e instanceof Error) {
                      setError(e.message);
                     }
              });
        
    };
    const handleMoreClick = () => {
        setCurrentPage((currentPage) => currentPage + 1);
    };
    return (
        <>
            <h1>Members</h1>
            {/* <pre>{JSON.stringify(MOCK_MEMBERS, null, ' ')}</pre> */}
            {error && (
                <div className="row">
                    <div className="card large error">
                        <section>
                            <p>
                                <span className="icon-alert inverse "></span>
                                {error}
                            </p>
                        </section>
                    </div>
                </div>
            )}
            <MemberList members={members} onSave={saveMember}/>
           {loading && (
                <div className="center-page">
                    <span className="spinner primary"></span>
                    <p>Loading...</p>
                </div>
           )}
            {!loading && !error && (
            <div className="row">
                <div className="col-sm-12">
                <div className="button-group fluid">
                    <button className="button default" onClick={handleMoreClick}>
                        More...
                    </button>
                </div>
                </div>
            </div>
            )}
        </>
    );
}

export default MembersPage;