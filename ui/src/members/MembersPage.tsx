import React from 'react';
import { useMembers } from './MemberHooks';
import MemberList from "../../src/members/MemberList"

function MembersPage() {
    const {
      members,
      loading,
      error,
      setCurrentPage,
      saveMember,
      saving,
      savingError,
    } = useMembers();

    const handleMoreClick = () => {
        setCurrentPage((currentPage) => currentPage + 1);
    };
    return (
        <>
            <h1>Members</h1>
            {/* <pre>{JSON.stringify(MOCK_MEMBERS, null, ' ')}</pre> */}
            {saving && <span className="toast">Saving...</span>}
            {(error || savingError) && (
                <div className="row">
                    <div className="card large error">
                        <section>
                            <p>
                                <span className="icon-alert inverse "></span>
                                {error} {savingError}
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