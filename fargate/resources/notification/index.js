var AWS = require('aws-sdk');

exports.handler = function(event, context) {
    var eventText = JSON.stringify(event, null, 2);
    console.log("Received event:", eventText);

    var job = event["CodePipeline.job"],
        jobId = job.id,
        params = JSON.parse(job.data.actionConfiguration.configuration.UserParameters);

    // Validate parameters passed in UserParameters
    if(! params || params.EnvName == '') {
        failure('The UserParameters is invalid');
        return;
    }

    // Publish to a SNS topic
    var sns = new AWS.SNS();
    sns.publish({
        TopicArn: process.env.TOPIC_ARN,
        Subject: "[ "+params.EnvName+" ] A new version was deployed via AWS CodePipeline",
        Message: "Successfully deployed a new version to "+params.EnvName+" fargate!"
    }, function (err, data) {
        if (err) {
            failure(err);
            return;
        }
        success("Function Finished.");
    });

    function success(message) {
        var api = new AWS.CodePipeline();
        api.putJobSuccessResult({jobId: jobId}, function (err, data) {
            if (err) {
                context.fail(err);
            } else {
                context.succeed(message);
            }
        });
    }
    function failure(message) {
        var api = new AWS.CodePipeline(),
            params = {
                jobId: jobId,
                failureDetails: {
                    type: 'JobFailed',
                    message: JSON.stringify(message),
                    externalExecutionId: context.invokeid
                }
            };
        api.putJobFailureResult(params, function (err, data) {
            context.fail(message);
        });
    }
};
