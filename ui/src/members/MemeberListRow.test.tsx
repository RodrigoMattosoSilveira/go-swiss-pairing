import { render, screen } from "@testing-library/react";
import React from "react";
import userEvent from '@testing-library/user-event';
import { MemoryRouter } from 'react-router-dom';
import { Member } from "./Member";
import MemberListRow from "./MemberListRow";
import renderer from "react-test-renderer";

describe("Member List Row", () => {
    let member: Member;
    let handleEdit: jest.Mock;

    const setup = () =>
        render(
            <MemoryRouter>
                <MemberListRow member={member} onEdit={handleEdit} />
            </MemoryRouter>
        );
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
        setup();
    });
    it("Member List Row member={member} onEdit={handleEdit}", () => {
        // Adding the Router is a hack to make this test work
        const tree = renderer
            .create(
                <MemoryRouter>
                    <MemberListRow member={member} onEdit={handleEdit} />
                </MemoryRouter>)
            .toJSON();
        expect(tree).toMatchSnapshot();

    });
    it('renders project properly', () => {
        setup();
        expect(screen.getByRole('generic', { name: 'member.email' })).toHaveTextContent(member.email);
        // screen.debug(document);
        screen.getByText(/Arianna.Moody@yahoo\.com/i);
        screen.getByText(/1026/i);
    });

    it('handler called when edit clicked', async () => {
        setup();
        // this query works screen.getByText(/edit/i)
        // but using role is better
        const user = userEvent.setup();
        await user.click(screen.getByRole('button', { name: 'member.id' }));
        expect(handleEdit).toBeCalledTimes(1);
        expect(handleEdit).toBeCalledWith(member);
   });
});