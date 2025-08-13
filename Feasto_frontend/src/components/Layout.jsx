import { Outlet } from "react-router-dom";
import { Navbar } from "./Navbar";

export default function Layout({ page, user }) {
  return (
    <>
      <Navbar page={page} user={user} />
        <Outlet /> {/* This is where child routes will render */}
    </>
);
}