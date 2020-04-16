import {
  expect as expectCDK,
  matchTemplate,
  MatchStyle,
} from "@aws-cdk/assert";
import * as cdk from "@aws-cdk/core";
import Infra = require("../lib/mirei-tts-stack");

test("Empty Stack", () => {
  const app = new cdk.App();
  // WHEN
  const stack = new Infra.MireiTTSStack(app, "MyTestStack");
  // THEN
  expectCDK(stack).to(
    matchTemplate(
      {
        Resources: {},
      },
      MatchStyle.EXACT
    )
  );
});
