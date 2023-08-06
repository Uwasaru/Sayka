type TColumn = {
  description: string;
  date: string;
};

// TODO: 本来はAPIから取得する
const columns: TColumn[] = [
  {
    description:
      "開発を本格的に開始しました。β版のリリースは2023年8月16日を予定しています。",
    date: "2023/08/06",
  },
];

export const Newses = () => {
  return (
    <div>
      <div className="my-2 text-lg font-bold">ニュース</div>
      <div className="space-y-5 text-sm">
        {columns.map((column) => (
          <News key={column.description} {...column} />
        ))}
      </div>
    </div>
  );
};

const News = (props: TColumn) => {
  return (
    <div className="space-y-2">
      <div>{props.description}</div>
      <div className="font-light opacity-80">{props.date}</div>
    </div>
  );
};
