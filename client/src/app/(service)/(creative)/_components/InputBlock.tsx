"use client";

import { ErrorMessage } from "@hookform/error-message";
import { FC } from "react";
import { useFormContext, RegisterOptions } from "react-hook-form";
import { AiFillGithub } from "react-icons/ai";

type TProps = {
  text: string;
  subText?: string;
  isRequired?: boolean;
  type?: string;
  name: string;
  options?: RegisterOptions;
  placeholder?: string;
  defaultValue?: string | number;
  icon?: React.ReactNode;
  feature?: "input" | "textarea";
  onChange?: (event: React.ChangeEvent<HTMLInputElement>) => void;
};

export const InputBlock: FC<TProps> = ({
  text,
  subText,
  isRequired = false,
  type = "text",
  name,
  options,
  placeholder,
  defaultValue,
  icon,
  feature = "input",
  onChange,
}) => {
  const {
    register,
    formState: { errors },
  } = useFormContext();

  const colorStyle = errors[name]
    ? "bg-red-100 focus:outline-red-200"
    : "border-gray-600 focus:bg-indigo-100 focus:outline-indigo-200";

  return (
    <div className="mb-5">
      {icon ? (
        <label htmlFor={name} className="mb-1">
          <div className="flex space-x-1">
            {icon}
            <span className="text-sm font-semibold">
              {text}
              {isRequired && <span className="text-red-600">*</span>}
            </span>
          </div>
          {subText && <div className="text-xs text-gray-400">{subText}</div>}
        </label>
      ) : (
        <label htmlFor={name} className="mb-1">
          <span className="text-sm font-semibold">
            {text}
            {isRequired && <span className="text-red-600">*</span>}
          </span>
          {subText && <div className="text-xs text-gray-400">{subText}</div>}
        </label>
      )}
      <div className="my-1 text-xs text-red-600">
        <ErrorMessage
          errors={errors}
          name={name}
          render={({ message }) => <p>{message}</p>}
        />
      </div>
      {feature === "textarea" ? (
        <textarea
          className={`${colorStyle} ${"w-full rounded border p-3 text-sm"}`}
          id={name}
          {...register(name, options)}
          placeholder={placeholder}
          defaultValue={defaultValue}
        />
      ) : (
        <input
          className={`${colorStyle} ${"w-full rounded border p-3 text-sm"}`}
          id={name}
          type={type}
          {...register(name, options)}
          placeholder={placeholder}
          defaultValue={defaultValue}
          onChange={onChange}
        />
      )}
    </div>
  );
};
