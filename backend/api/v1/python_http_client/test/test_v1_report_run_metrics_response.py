# coding: utf-8

"""
    Kubeflow Pipelines on Tekton API

    This file contains REST API specification for Kubeflow Pipelines on Tekton. The file is autogenerated from the swagger definition.

    Contact: prashsh1@in.ibm.com
    Generated by: https://openapi-generator.tech
"""


from __future__ import absolute_import

import unittest
import datetime

import kfp_tekton_server_api
from kfp_tekton_server_api.models.v1_report_run_metrics_response import V1ReportRunMetricsResponse  # noqa: E501
from kfp_tekton_server_api.rest import ApiException

class TestV1ReportRunMetricsResponse(unittest.TestCase):
    """V1ReportRunMetricsResponse unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test V1ReportRunMetricsResponse
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # model = kfp_tekton_server_api.models.v1_report_run_metrics_response.V1ReportRunMetricsResponse()  # noqa: E501
        if include_optional :
            return V1ReportRunMetricsResponse(
                results = [
                    kfp_tekton_server_api.models.report_run_metrics_response_report_run_metric_result.ReportRunMetricsResponseReportRunMetricResult(
                        metric_name = '0', 
                        metric_node_id = '0', 
                        status = 'UNSPECIFIED', 
                        message = '0', )
                    ]
            )
        else :
            return V1ReportRunMetricsResponse(
        )

    def testV1ReportRunMetricsResponse(self):
        """Test V1ReportRunMetricsResponse"""
        inst_req_only = self.make_instance(include_optional=False)
        inst_req_and_optional = self.make_instance(include_optional=True)


if __name__ == '__main__':
    unittest.main()
