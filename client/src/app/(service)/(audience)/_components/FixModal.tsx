"use client";

import { useToast } from "@chakra-ui/react";
import Link from "next/link";
import { useRouter } from "next/navigation";
import React, { FC, useRef, useEffect, useState } from "react";
import { AiOutlineDelete, AiOutlineEdit } from "react-icons/ai";

import { deleteSayka } from "@/api";

type TProps = {
  isVisible: boolean;
  onClose: () => void;
  id: string;
  token?: string;
};

export const FixModal: FC<TProps> = ({ isVisible, onClose, id, token }) => {
  const [isConfirmDeleteVisible, setConfirmDeleteVisible] = useState(false);
  const modalRef = useRef<HTMLDivElement>(null);
  const r = useRouter();
  const toast = useToast();

  const handleCloseModal = (e: MouseEvent) => {
    if (modalRef.current && !modalRef.current.contains(e.target as Node)) {
      onClose();
      setConfirmDeleteVisible(false);
    }
  };

  useEffect(() => {
    document.addEventListener("mousedown", handleCloseModal);
    return () => {
      document.removeEventListener("mousedown", handleCloseModal);
    };
  });

  const handleDelete = async () => {
    if (!token) return;
    await deleteSayka(id, token);
    setConfirmDeleteVisible(false); // 確認モーダルを閉じる
    onClose(); // メインモーダルも閉じる
    r.refresh();
    toast({
      position: "bottom-left",
      render: () => (
        <div className="bg-gray-800 p-3 text-white">Saykaを削除しました。</div>
      ),
    });
  };

  const handleCancel = () => {
    setConfirmDeleteVisible(false);
    onClose();
  };

  if (!isVisible) return null;

  return (
    <div
      ref={modalRef}
      className="absolute right-0 top-0 z-10 w-48 rounded border border-gray-200 bg-white shadow-md"
      onClick={(e) => e.stopPropagation()}>
      {isConfirmDeleteVisible ? (
        <div>
          <button
            onClick={handleDelete}
            className="flex w-full items-center px-4 py-2 text-left font-semibold text-red-800 hover:bg-gray-100">
            本当に削除する
          </button>
          <button
            onClick={handleCancel}
            className=" flex w-full items-center px-4 py-2 text-left font-semibold hover:bg-gray-100">
            キャンセル
          </button>
        </div>
      ) : (
        <div>
          <Link
            href={`/edit/${id}`}
            className="flex w-full items-center px-4 py-2 text-left font-semibold text-green-700 hover:bg-gray-100">
            <AiOutlineEdit size={20} className="mr-2" />
            編集
          </Link>
          <button
            onClick={() => setConfirmDeleteVisible(true)}
            className=" flex w-full items-center px-4 py-2 text-left font-semibold text-red-800 hover:bg-gray-100">
            <AiOutlineDelete size={20} className="mr-2" />
            削除
          </button>
        </div>
      )}
    </div>
  );
};
