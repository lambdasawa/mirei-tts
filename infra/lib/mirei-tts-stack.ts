import * as path from "path";
import * as cdk from "@aws-cdk/core";
import * as s3 from "@aws-cdk/aws-s3";
import * as ec2 from "@aws-cdk/aws-ec2";
import * as ecr from "@aws-cdk/aws-ecr";
import * as ecsAssets from "@aws-cdk/aws-ecr-assets";
import * as ecs from "@aws-cdk/aws-ecs";
import * as ecsPatterns from "@aws-cdk/aws-ecs-patterns";
import * as route53 from "@aws-cdk/aws-route53";
import * as route53Targets from "@aws-cdk/aws-route53-targets";

export class MireiTTSStack extends cdk.Stack {
  readonly port = 3011;

  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    const asset = new ecsAssets.DockerImageAsset(this, "DockerImageAsset", {
      directory: path.join(__dirname, "../../"),
      exclude: ["cdk.out", "node_modules"],
    });

    const vpc = ec2.Vpc.fromLookup(this, "VPC", {
      isDefault: true,
    });

    const cluster = new ecs.Cluster(this, "Cluster", {
      vpc: vpc,
    });

    new ecsPatterns.ApplicationLoadBalancedFargateService(this, "Service", {
      cluster,
      memoryLimitMiB: 4096,
      cpu: 2048,
      assignPublicIp: true,
      taskImageOptions: {
        image: ecs.ContainerImage.fromEcrRepository(
          asset.repository,
          asset.sourceHash
        ),
        containerPort: this.port,
        environment: this.getEnvironment(),
      },
    });
  }

  getEnvironment(): Record<string, string> {
    const environment: Record<string, string> = {};

    for (const key in process.env) {
      if (!key.startsWith("MTTS_")) continue;

      environment[key] = process.env[key] || "";
    }

    environment["MTTS_ADDRESS"] = `:${this.port}`;

    return environment;
  }
}
