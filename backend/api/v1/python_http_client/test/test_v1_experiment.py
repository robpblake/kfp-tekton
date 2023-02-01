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
from kfp_tekton_server_api.models.v1_experiment import V1Experiment  # noqa: E501
from kfp_tekton_server_api.rest import ApiException

class TestV1Experiment(unittest.TestCase):
    """V1Experiment unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test V1Experiment
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # model = kfp_tekton_server_api.models.v1_experiment.V1Experiment()  # noqa: E501
        if include_optional :
            return V1Experiment(
                id = '0', 
                name = '0', 
                description = '0', 
                created_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                resource_references = [
                    kfp_tekton_server_api.models.v1_resource_reference.v1ResourceReference(
                        key = kfp_tekton_server_api.models.v1_resource_key.v1ResourceKey(
                            type = 'UNKNOWN_RESOURCE_TYPE', 
                            id = '0', ), 
                        name = '0', 
                        relationship = 'UNKNOWN_RELATIONSHIP', )
                    ], 
                storage_state = 'STORAGESTATE_UNSPECIFIED'
            )
        else :
            return V1Experiment(
        )

    def testV1Experiment(self):
        """Test V1Experiment"""
        inst_req_only = self.make_instance(include_optional=False)
        inst_req_and_optional = self.make_instance(include_optional=True)


if __name__ == '__main__':
    unittest.main()
