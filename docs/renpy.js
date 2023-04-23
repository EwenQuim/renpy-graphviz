// import { request } from "https://cdn.skypack.dev/@octokit/request";
//contains the graphviz file as string

function printMessagePromise(msg, boolChoice1, boolChoice2) {
  return new Promise((resolve, reject) => {
    printMessage(msg, boolChoice1, boolChoice2, (err, message) => {
      console.log("cb", msg, err);
      if (err) {
        reject(err);
        return;
      }
      resolve(message);
    });
  });
}

async function getRenpy(repoName, subPath) {
  var errorMessage = document.getElementById("errorMessage");
  errorMessage.style.visibility = "hidden";

  var mainResponse;
  var renpyString = "";
  console.log("fetching start");
  if (subPath) {
    mainResponse = await fetch(
      "https://api.github.com/search/code?accept=application/vnd.github.v3+json&q=label+path:" +
        subPath +
        "+extension:rpy+repo:" +
        repoName
    );
  } else {
    mainResponse = await fetch(
      "https://api.github.com/search/code?accept=application/vnd.github.v3+json&q=label+extension:rpy+repo:" +
        repoName
    );
  }

  // const result = await request("GET /search/code", {
  //   headers: {
  //     authorization: "token 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",
  //   },
  //   q: "label+repo:qirien/personal-space+extension:rpy",
  // });
  // console.log(result);

  if (mainResponse.status != 200) {
    errorMessage.style.visibility = "visible";
  }
  const mainAns = await mainResponse.json();

  console.log("fetching end", mainResponse, mainAns);

  const ttt = await Promise.all(
    mainAns.items
      .filter(
        (item) =>
          !item.path.includes("tl/") &&
          !item.path.includes("options.rpy") &&
          !item.path.includes("gui.rpy") &&
          !item.path.includes("screens.rpy") &&
          !item.path.includes("00")
      )
      .map(async (item) => {
        const rawFileUrl = item.html_url
          .replace("github.com", "raw.githubusercontent.com")
          .replace("blob/", "");
        // console.log(rawFileUrl);
        const rep = await fetch(rawFileUrl);
        return await rep.text();
      })
  );

  for (const ans of ttt) {
    // console.log(item.path);

    renpyString = renpyString.concat(ans).concat("\n#renpy-graphviz: BREAK\n");
  }

  return renpyString;
}

function getRepoStruct(s) {
  const regex = /\w*\/\w*\//g;

  if (regex.test(s)) {
    const last = regex.lastIndex;
    console.log(s.substring(0, last), s.substring(last));

    return [s.substring(0, last), s.substring(last)];
  } else {
    console.log(s, null);
    return [s, null];
  }
}

export async function getRepo() {
  const loader = document.getElementById("loader");
  loader.style.visibility = "visible";

  var [repoName, subPath] = getRepoStruct(
    document.getElementById("repo").value.trim()
  );
  console.log(repoName, subPath);
  const renpyTextList = await getRenpy(repoName, subPath);
  const graph = await printMessagePromise(
    renpyTextList,
    document.getElementById("choices").checked,
    !document.getElementById("hideAtoms").checked
  );

  console.log(graph);

  var t = d3.transition().duration(750).ease(d3.easeLinear);

  d3.select("#graph").graphviz().transition(t).renderDot(graph);
  loader.style.visibility = "hidden";
}
