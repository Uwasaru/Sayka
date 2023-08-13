// Modal.tsx
import Link from "next/link";
import React, { FC, useRef, useEffect } from "react";
import { AiOutlineDelete, AiOutlineEdit } from "react-icons/ai";

type TProps = {
  isVisible: boolean;
  onClose: () => void;
  id: number;
};

export const FixModal: FC<TProps> = ({ isVisible, onClose, id }) => {
  const modalRef = useRef<HTMLDivElement>(null);

  const handleCloseModal = (e: MouseEvent) => {
    if (modalRef.current && !modalRef.current.contains(e.target as Node)) {
      onClose();
    }
  };

  useEffect(() => {
    document.addEventListener("mousedown", handleCloseModal);
    return () => {
      document.removeEventListener("mousedown", handleCloseModal);
    };
  });

  if (!isVisible) return null;

  return (
    <div
      ref={modalRef}
      className="absolute right-0 top-0 z-10 w-48 border border-gray-200 bg-white shadow-md rounded"
      onClick={(e) => e.stopPropagation()}>
      <Link
        href={`/edit/${id}`}
        className="w-full px-4 py-2 text-left hover:bg-gray-100 font-semibold flex items-center text-green-700">
        <AiOutlineEdit size={20} className="mr-2" />
        編集
      </Link>
      <button
        onClick={() => console.log("削除")}
        className=" w-full px-4 py-2 text-left hover:bg-gray-100 font-semibold flex items-center text-red-800">
        <AiOutlineDelete size={20} className="mr-2" />
        削除
      </button>
    </div>
  );
};