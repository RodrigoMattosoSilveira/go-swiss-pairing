import { render, screen } from "@testing-library/react";
import React from "react";
import { Member } from "./Member";
import MemberListRow from "./MemberListRow";

describe("<MemberCard />", () => {
    let member: Member;
    let handleEdit: jest.Mock;
    beforeEach(() => {
        member = new Member({
            "id":"mkq9OLj-Wf",
            "first":"Arianna",
            "last":"Moody",
            "email":"Arianna.Moody@yahoo.com",
            "password":"VH8q4tkg6UX3",
            "cell":"801 277-6891",
            "rating":1026,
            "isActive":true,
            "imageUrl":"/assets/picture_mkq9OLj-Wf.jpeg"},);
        handleEdit = jest.fn();
    });

    it("should initially render", () => {
        render(<MemberListRow member={member} onEdit={handleEdit} />);
    });
});